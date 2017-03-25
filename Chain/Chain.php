<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 7:16
 */

abstract class AbstractLogger{
    const INFO=1;
    const DEBUG=2;
    const ERROR=3;
    protected $level;

    protected $nextLogger;

    public function setNextLogger($nextLogger){
        $this->nextLogger=$nextLogger;
    }

    public function logMessage($level,$message){
        if($this->level<=$level){
            $this->write($message);
        }

        if($this->nextLogger!=null){
            $this->nextLogger->logMessage($level,$message);
        }
    }

    abstract protected function write($message);
}


class ConsoleLogger extends AbstractLogger{
    public function __construct($level)
    {
        $this->level=$level;
    }

    public function write($message)
    {
        // TODO: Implement write() method.
        var_dump("Standard Console::Logger: ".$message);
    }
}



class ErrorLogger extends AbstractLogger{
    public function __construct($level)
    {
        $this->level=$level;
    }

    public function write($message)
    {
        // TODO: Implement write() method.
        var_dump("Error Console::Logger: ".$message);
    }
}

class FileLogger extends AbstractLogger{
    public function __construct($level)
    {
        $this->level=$level;
    }

    public function write($message)
    {
        // TODO: Implement write() method.
        var_dump("File::Logger: ".$message);
    }
}

function getChainOfLoggers() {
    $errorLogger=new ErrorLogger(AbstractLogger::ERROR);
    $fileLogger=new FileLogger(AbstractLogger::DEBUG);
    $consoleLogger=new ConsoleLogger(AbstractLogger::INFO);
    $errorLogger->setNextLogger($fileLogger);
    $fileLogger->setNextLogger($consoleLogger);
    return $errorLogger;
}

//demo
$logchain=getChainOfLoggers();
$logchain->logMessage(AbstractLogger::INFO,"This is an information");
$logchain->logMessage(AbstractLogger::DEBUG,"This is an debug level information");
$logchain->logMessage(AbstractLogger::ERROR,"This is an error information");