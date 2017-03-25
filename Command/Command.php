<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 9:14
 */

class Stock{
    private $name="ABC",$quantity=10;

    public function buy(){
      var_dump("stock [name".$this->name.",quantity:".$this->quantity."] bought");
    }

    public function sell(){
        var_dump("stock [name".$this->name.",quantity:".$this->quantity."] sold");
    }
}

interface Order{
    public function execute();
}

class BuyOrder implements Order{
    private $stock;
    public function __construct($stock)
    {
        $this->stock=$stock;
    }
    public function execute(){
        $this->stock->buy();
    }
}


class SellOrder implements Order{
    private $stock;
    public function __construct($stock)
    {
        $this->stock=$stock;
    }
    public function execute(){
        $this->stock->sell();
    }
}

class Broker{
    private $orderList=array();
    public function takeOrder($order){
        $this->orderList[]=$order;
    }

    public function placeOrder(){
        foreach($this->orderList as $order){
            $order->execute();
        }
        $this->orderList=array();
    }
}


//demo
$stock=new Stock();
$buyorder=new BuyOrder($stock);
$sellorder=new SellOrder($stock);
$broker=new Broker();
$broker->takeOrder($buyorder);
$broker->takeOrder($sellorder);
$broker->takeOrder($buyorder);
$broker->placeOrder();
