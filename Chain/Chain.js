/**
 * Created by jamesyan on 2017/3/25.
 */
const ERROR = 3;
const DEBUG = 2;
const INFO = 1;

var AbstractLogger = function (level) {
    this.level = level;
    this.write=function(){};
}

AbstractLogger.prototype.setNextLogger = function (nextLogger) {
    this.nextLogger = nextLogger;
}


AbstractLogger.prototype.logMessage = function (level, message) {
    if (this.level <= level) {
        this.write.call(this,message);
    }

    if (this.nextLogger != null) {
        this.nextLogger.logMessage(level, message);
    }
}

var ErrorLogger = function (level) {
    AbstractLogger.call(this, level);
    this.write = function (message) {
        console.log("Error Logger:" + message);
    }
}

ErrorLogger.prototype = Object.create(AbstractLogger.prototype);

var FileLogger = function (level) {
    AbstractLogger.call(this, level);
    this.write = function (message) {
        console.log("File Debug Logger:" + message);
    }
}

FileLogger.prototype = Object.create(AbstractLogger.prototype);

var ConsoleLogger = function (level) {
    AbstractLogger.call(this, level);
    this.write = function (message) {
        console.log("standard Console Logger:" + message);
    }
}

ConsoleLogger.prototype = Object.create(AbstractLogger.prototype);

function chainPatternDemo() {
    var errorLogger = new ErrorLogger(ERROR);
    var fileLogger = new FileLogger(DEBUG);
    var consoleLogger = new ConsoleLogger(INFO);
    errorLogger.setNextLogger(fileLogger);
    fileLogger.setNextLogger(consoleLogger)
    return errorLogger;

}

