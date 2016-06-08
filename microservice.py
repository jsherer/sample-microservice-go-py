import flask
app = flask.Flask(__name__)

@app.route('/hello')
def hello():
    return 'Hello!'

@app.route('/sleep')
def sleep():
    import time
    time.sleep(10)
    return 'Done!'
