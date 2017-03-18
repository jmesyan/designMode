/**
 * Created by jamesyan on 2017/3/15.
 */

/**
interface Item
{
    public function name();

    public function packing();

    public function price();
}

interface Packing
{
    public function Pack();
}

 */


var Packing={
    Wrapper:function(){
        return "Wrapper";
    },
    Bottle:function(){
        return "Bottle";
    }
}

var Item={
    Burger:{
        packing:function(){
            return Packing.Wrapper();
        },
        vegBurger:function(){
            return {
                pack: function () {
                    return Item.Burger.packing()
                },
                name: function () {
                    return "veg Burger";
                },
                price: function () {
                    return 25.3;
                }
            }
        },
        chickenBurger:function(){
            return {
                pack: function () {
                    return Item.Burger.packing()
                },
                name: function () {
                    return "chicken Burger";
                },
                price: function () {
                    return 36.3;
                }
            }
        },
    },
    ColdDrink:{
        packing:function(){
            return Packing.Bottle();
        },
        Coke:function(){
            return {
                pack: function () {
                    return Item.ColdDrink.packing()
                },
                name: function () {
                    return "coke";
                },
                price: function () {
                    return 28.3;
                }
            }
        },
        pesi:function(){
            return {
                pack: function () {
                    return Item.ColdDrink.packing()
                },
                name: function () {
                    return "pesi";
                },
                price: function () {
                    return 15.6;
                }
            }
        }
    }

}

var Meal=function(){
    return {
        items: [],
        addItems: function (item) {
            this.items.push(item);
        },
        showItems: function () {
            var show='';
            for (var i = 0; i < this.items.length; i++) {
                var item = this.items[i];
                show+="item: " + item.name() + ",Packing :" + item.pack() + ",Price :" + item.price()+"\r\n";
            }
            return show;
        },
        getCost: function () {
            var cost = 0;
            for (var i = 0; i < this.items.length; i++) {
                var item = this.items[i];
                cost += item.price();
            }
            return cost;
        }
    }
}

var MealManger=function(){
    this.prepareVegMeals=function(){
        var meal=new Meal();
        meal.addItems(new Item.Burger.vegBurger);
        meal.addItems(new Item.ColdDrink.Coke);
        return meal;
    };

    this.prepareNonVegMeals=function(){
        var meal=new Meal();
        meal.addItems(new Item.Burger.chickenBurger);
        meal.addItems(new Item.ColdDrink.pesi);
        return meal;
    }
}