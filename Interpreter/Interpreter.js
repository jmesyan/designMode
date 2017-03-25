/**
 * Created by jamesyan on 2017/3/25.
 */

var TerminalExpress=function(data){
    this.data=data;
    this.interpret=function(context){
        if(context.indexOf(this.data)!=-1){
            return true;
        }
        return false;
    }
}

var AndExpress=function(exp1,exp2){
    this.exp1=exp1;
    this.exp2=exp2;
    this.interpret=function(context){
        return this.exp1.interpret(context) &&this.exp2.interpret(context);
    }
}

var OrExpress=function(exp1,exp2){
    this.exp1=exp1;
    this.exp2=exp2;
    this.interpret=function(context){
        return this.exp1.interpret(context) ||this.exp2.interpret(context);
    }
}


function isMale(){
    var robert=new TerminalExpress("Robert");
    var join=new TerminalExpress("Join");
    return new OrExpress(robert,join);
}


function isMarriedWoman(){
    var married=new TerminalExpress("Married");
    var juliy=new TerminalExpress("Julie");
    return new AndExpress(married,juliy);
}