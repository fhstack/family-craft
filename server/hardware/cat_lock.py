import RPi.GPIO as GPIO
import time
import sys

P_SERVO = 8  # GPIO端口号，根据实际修改
fPWM = 50  # Hz (软件PWM方式，频率不能设置过高)
b = 2.5
GPIO.setmode(GPIO.BOARD)
GPIO.setup(P_SERVO, GPIO.OUT)
pwm = GPIO.PWM(P_SERVO, fPWM)
pwm.start(0)

def set_direction(direction):
    duty = (1 / 27) * direction + b
    pwm.ChangeDutyCycle(duty)
    print("direction =", direction, "-> duty =", duty)

angle = sys.argv[1]
set_direction(int(angle))
time.sleep(0.5)
pwm.stop()
GPIO.cleanup()
