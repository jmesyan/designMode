<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 6:46
 */

interface Image{
    public function display();
}


class RealImage implements Image {
    private $fileName;
    public function __construct($fileName)
    {
        $this->fileName=$fileName;
        $this->loadFromDisk($fileName);
    }

    public function display(){
        echo "displaying :".$this->fileName."<br/>";
    }
    public function loadFromDisk($fileName){
        echo "Loding disk".$fileName."<br/>";
    }
}

class ProxyImage implements Image{
    private $realImage,$fileName;
    public function __construct($fileName)
    {
        $this->fileName = $fileName;
    }

    public function display()
    {
        // TODO: Implement display() method.
        if(!$this->realImage instanceof RealImage){
            $this->realImage=new RealImage($this->fileName);
        }
        $this->realImage->display();
    }
}


//demo
//图像从磁盘加载
$image=new ProxyImage("abcx.jpg");
$image->display();
//图像将无法从磁盘加载
$image->display();


