import RPi.GPIO as GPIO
import time
from http.server import HTTPServer, BaseHTTPRequestHandler

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
    time.sleep(0.04)
    pwm.ChangeDutyCycle(0)


host = ('localhost', 8888)


class Resquest(BaseHTTPRequestHandler):
    def do_GET(self):
        angle = self.path[1:]
        print(angle)
        set_direction(int(angle))
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()


if __name__ == '__main__':
    server = HTTPServer(host, Resquest)
    print("Starting server, listen at: %s:%s" % host)
    server.serve_forever()

pwm.stop()
GPIO.cleanup()
