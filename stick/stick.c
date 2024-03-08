#include <linux/joystick.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <stdio.h>
#include <stdbool.h>
#include <assert.h>
#include <string.h>

#define JOYSTICK_MAX 4
#define JOYSTICK_AXIS_MAX 16
#define JOYSTICK_BUTTON_MAX 32

#define JOY_DEV "/dev/input/js"

// typedef enum JOY_BUTTONS
// {
// 	JOY_BUTTON_X,
// 	JOY_BUTTON_A,
// 	JOY_BUTTON_B,
// 	JOY_BUTTON_Y,
// 	JOY_BUTTON_LEFTSHOULDER,
// 	JOY_BUTTON_RIGHTSHOULDER,
// 	JOY_BUTTON_LEFTTRIGGER,
// 	JOY_BUTTON_RIGHTTRIGGER,
// 	JOY_BUTTON_BACK,
// 	JOY_BUTTON_START,
// 	JOY_BUTTON_LEFTSTICK,
// 	JOY_BUTTON_RIGHTSTICK,
// 	JOY_BUTTONS_DEFINED,
// } JOY_BUTTONS;

static const char *button_labels[JOYSTICK_BUTTON_MAX] = {
	"x",
	"a",
	"b",
	"y",
	"leftshoulder",
	"rightshoulder",
	"lefttrigger",
	"righttrigger",
	"back",
	"start",
	"leftstick",
	"rightstick",
};

static char label_buffer[16];

static const char *button_label(size_t i)
{
	if (i < 0 || i >= JOYSTICK_BUTTON_MAX)
	{
		snprintf(label_buffer, sizeof(label_buffer), "button%d", (int)i);
		return label_buffer;
	}
	return button_labels[i];
}

static const char *axis_labels[JOYSTICK_AXIS_MAX] = {};

typedef struct joy_stick
{
	int handle;
	int axis_count;
	int button_count;
	char name[80];

	u_int64_t button_state;
	u_int64_t button_prev;

	int16_t axis[JOYSTICK_AXIS_MAX];
	int16_t axis_prev[JOYSTICK_AXIS_MAX];

	struct js_corr corr;
	__u16 btn_map[KEY_MAX - BTN_MISC + 1];
	__u8 axis_map[ABS_CNT];
} joy_stick;

static bool initialized = false;
static joy_stick joy_sticks[JOYSTICK_MAX] = {0};
static int joy_stick_count = 0;
static struct js_event last_button_pressed_event = {0};

static const struct js_event zero_event = {0};

#define event_size sizeof(struct js_event)

#define in_range(a) ((a >= 0 && a < joy_stick_count))
#define not_in_range(a) ((a < 0 || a >= joy_stick_count))

#define is_axis_event(e) ((e.type & ~JS_EVENT_INIT) == JS_EVENT_AXIS)
#define is_this_axis(e, a) ((e.number == a))
// #define axis_event_index(e) (e.number)
// #define axis_event_value(e) (e.value)

#define is_button_event(e) ((e.type & ~JS_EVENT_INIT) == JS_EVENT_BUTTON)
#define button_event_index(e) (e.number & 0x3f)
#define is_button_event_down(e) ((e.value == 1))
#define is_button_event_up(e) ((e.value == 0))

#define set_button_state_down(s, b) ((s |= (1 << b)))
#define set_button_state_up(s, b) ((s &= ~(1 << b)))
#define is_button_state_down(s, b) ((s & (1 << b)))
#define is_button_state_up(s, b) (!(s & (1 << b)))

void initialize()
{
	char device[80] = {0};

	joy_stick_count = 0;

	for (int i = 0; i < JOYSTICK_MAX; i++)
	{
		joy_stick *joy = &joy_sticks[joy_stick_count];
		snprintf(device, sizeof(device), JOY_DEV "%d", i);
		joy->handle = open(device, O_RDONLY);
		if (joy->handle == -1)
		{
			continue;
		}
		ioctl(joy->handle, JSIOCGAXES, &joy->axis_count);
		ioctl(joy->handle, JSIOCGBUTTONS, &joy->button_count);
		ioctl(joy->handle, JSIOCGNAME(sizeof(joy->name)), &joy->name);

		ioctl(joy->handle, JSIOCGCORR, &joy->corr);
		ioctl(joy->handle, JSIOCGBTNMAP, joy->btn_map);
		ioctl(joy->handle, JSIOCGAXMAP, joy->axis_map);

		// Set nonblocking
		fcntl(joy->handle, F_SETFL, O_NONBLOCK);
		++joy_stick_count;
	}

	initialized = true;
}

__u8 axis_custom[4][8] = {
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
};

__u8 axis_default[4][8] = {
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
	{0x00, 0x01, 0x02, 0x05, 0x16, 0x17, 0x00, 0x00},
};

bool UpdateAxisMap(int Joystick, __u8 *axis_map)
{
	if (not_in_range(Joystick))
	{
		return false;
	}

	// JSIOCSAXMAP
	joy_stick *joy = &joy_sticks[Joystick];
	ioctl(joy->handle, JSIOCGAXMAP, joy->axis_map);
	return false;
}

bool IsJoystickAvailable(int Joystick)
{
	return in_range(Joystick);
}

const char *GetJoystickName(int Joystick)
{
	if (not_in_range(Joystick))
	{
		return "undefined";
	}
	return joy_sticks[Joystick].name;
}

const char *GetButtonName(int Joystick, int button)
{
	// if (not_in_range(Joystick))
	// {
	// 	return "undefined";
	// }

	return button_label(button);
}

bool IsJoystickButtonPressed(int Joystick, int button)
{
	if (not_in_range(Joystick))
	{
		return false;
	}

	joy_stick *joy = &joy_sticks[Joystick];
	if (button >= joy->button_count)
	{
		return false;
	}

	return is_button_state_down(joy->button_prev, button) ||
		   is_button_state_down(joy->button_state, button);
}

bool IsJoystickButtonDown(int Joystick, int button)
{
	if (not_in_range(Joystick))
	{
		return false;
	}
	joy_stick *joy = &joy_sticks[Joystick];

	if (button >= joy->button_count)
	{
		return false;
	}

	return is_button_state_down(joy_sticks[Joystick].button_state, button);
}

bool IsJoystickButtonReleased(int Joystick, int button)
{
	if (not_in_range(Joystick))
	{
		return false;
	}

	joy_stick *joy = &joy_sticks[Joystick];
	if (button >= joy->button_count)
	{
		return false;
	}

	return is_button_state_down(joy->button_prev, button) &&
		   is_button_state_up(joy->button_state, button);
}

bool IsJoystickButtonUp(int Joystick, int button)
{
	if (not_in_range(Joystick))
	{
		return false;
	}
	joy_stick *joy = &joy_sticks[Joystick];

	if (button >= joy->button_count)
	{
		return false;
	}
	return is_button_state_up(joy->button_state, button);
}

int GetJoystickButtonPressed(void)
{
	return button_event_index(last_button_pressed_event);
}

int GetJoystickButtonCount(int Joystick)
{
	if (not_in_range(Joystick))
	{
		return 0;
	}
	return joy_sticks[Joystick].button_count;
}

int GetJoystickAxisCount(int Joystick)
{
	if (not_in_range(Joystick))
	{
		return 0;
	}
	return joy_sticks[Joystick].axis_count;
}

int16_t GetJoystickAxisValue(int Joystick, int axis)
{
	if (not_in_range(Joystick))
	{
		return 0;
	}

	joy_stick *joy = &joy_sticks[Joystick];
	if (axis < 0 || axis >= joy->axis_count)
	{
		return 0;
	}

	return joy->axis[axis];
}

float GetJoystickAxisMovement(int Joystick, int axis)
{
	// #define DPLEFT 0
	// #define DPDOWN 1
	// #define DPRIGHT 2
	// #define DPUP 3

	// 	static float dpi[4] = {.4, .8, .2, .1};

	if (not_in_range(Joystick))
	{
		return 0;
	}

	joy_stick *joy = &joy_sticks[Joystick];
	if (axis < 0 || axis >= joy->axis_count)
	{
		return 0;
	}

	return (float)(joy->axis[axis] - joy->axis_prev[axis]);
}

int SetJoystickMappings(const char *mappings)
{
	return 0;
}

void BeginJoystick()
{
	if (!initialized)
	{
		initialize();
	}

	struct js_event event;
	struct js_event button_pressed_event = last_button_pressed_event;

	joy_stick *joy;
	int button, down;

	for (int i = 0; i < joy_stick_count; i++)
	{
		joy = &joy_sticks[i];

		while (event_size == read(joy->handle, &event, event_size))
		{
			if (is_axis_event(event))
			{
				int index = event.number;
				joy->axis_prev[index] = joy->axis[index];
				joy->axis[index] = event.value;
			}
			else if (is_button_event(event))
			{
				joy->button_prev = joy->button_state;
				button = button_event_index(event);
				if (is_button_event_down(event))
				{
					set_button_state_down(joy->button_state, button);
				}
				else
				{
					// last button up
					button_pressed_event = event;
					set_button_state_up(joy->button_state, button);
				}
			}
		}
	}
	last_button_pressed_event = button_pressed_event;
}

void dump_event(struct js_event *p_event)
{
	printf("time: %xms, value: %x, type: %x, number: %x\n",
		   p_event->time,
		   p_event->value,
		   p_event->type,
		   p_event->number);
}

void dump_corr(struct js_corr *p_corr)
{
	printf("correction values:\n");
	printf("type:%d, prec:%d, coef: [%d:%d:%d:%d:%d:%d:%d:%d]\n",
		   p_corr->type, p_corr->prec,
		   p_corr->coef[0], p_corr->coef[1], p_corr->coef[2], p_corr->coef[3],
		   p_corr->coef[4], p_corr->coef[5], p_corr->coef[6], p_corr->coef[7]);
}

void dump_axis_map(__u8 *map)
{
	printf("axis map:\n");
	for (size_t i = 0; i < ABS_CNT; i += 16)
	{
		printf("[%d:%d:%d:%d:%d:%d:%d:%d]\n",
			   map[0], map[1], map[2], map[3],
			   map[4], map[5], map[6], map[7]);
	}
}

void dump_button_map(__u16 *map)
{
	printf("button map:\n");
	for (size_t i = 0; i < KEY_MAX - BTN_MISC + 1; i += 16)
	{
		printf("[%d:%d:%d:%d:%d:%d:%d:%d]\n",
			   map[0], map[1], map[2], map[3],
			   map[4], map[5], map[6], map[7]);
	}
}

void DumpJoystick(void)
{
	for (size_t i = 0; i < joy_stick_count; i++)
	{
		joy_stick *joy = &joy_sticks[i];
		printf("Joystick: %d, %s, %d axes, %d buttons\n",
			   (int)i, joy->name, joy->axis_count, joy->button_count);
		dump_corr(&joy->corr);
		dump_axis_map(joy->axis_map);
		dump_button_map(joy->btn_map);
	}
}
