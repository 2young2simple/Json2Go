var fs= require('fs');
var path = require('path');

//  go 语法 
var gStr = "string"
var gInt = "int"
var gBool = "bool"
var gArray = "[]string"
var gArrayO = "[]interface{}"
var gObject = "interface{}"
var pack = "package node \n"
var header = (name) => "type "+name+" struct { \n"
var end = "\n} \n"

fs.writeFile("json.go",pack,'utf8',function(err){  
    if(err){  
        console.log(err);
    }
});

//  检测类型
function check(check){
    var type = Object.prototype.toString.call(check)
    switch(type){
        case "[object Number]":return gInt;break;
        case "[object String]":return gStr;break;
        case "[object Object]":return gObject;break;
        case "[object Array]":return gArrayO;break;
        case "[object Boolean]":return gBool;break;
        default :return gObject;
    }
}


//  首字母大写
function toUpper(str){
    var arr = str.split("");
    arr[0] = arr[0].toUpperCase();
    return arr.join("");
}
//  go参数命名规则  驼峰
function toGoName(str){
    str = str.replace("-","_")
    var arr = str.split("_");
    for (var i = 0;i < arr.length;i++){
        arr[i] = toUpper(arr[i]);        
    }
    return arr.join("");
}

function toJsonOrmTag(key){
    key = key.replace("-","_").toLowerCase()
    return "`orm:\"column("+key+")\"  json:"+"\""+key+"\"`"
}
function toJsonTag(key){
    key = key.replace("-","_").toLowerCase()
    return "`json:"+"\""+key+"\"`"
}

function toTime(){
    var create_time = "CreateTime    time.Time `orm:\"column(create_time);type(datetime);auto_now_add\"    json:\"create_time\"` \n"
    var update_time = "UpdateTime    time.Time `orm:\"column(update_time);type(datetime);auto_now\"    json:\"update_time\"` \n"
    return create_time+update_time
}

//  json 转 go struct
function toStruct(name,json){
    if (check(json) == gStr){
        return gStr
    }
    if (check(json) == gInt){
        return gInt
    }
    var str = header(toGoName(name));

    var keys = [];
    var isEmptyObject = true
    for (var key in json){
        isEmptyObject = false
        var ty = check(json[key])
        if (gObject == check(json[key])){
            ty = toStruct(key,json[key])
        }
        if (gArrayO == check(json[key])){
            ty = "[]"+toStruct(key,json[key][0])
        }
        str = str+toGoName(key)+" "+ty+" "+toJsonTag(key)+"\n";
    }
    if(isEmptyObject){
        return gObject
    }
    // orm create_time and update_time
    // str += toTime();
    str += end;
    
    fs.appendFile("json.go",str,'utf8',function(err){  
        if(err){  
            console.log(err);
        }
    });

    return toGoName(name)
}


var file = fs.readFileSync("source.json", "utf8");
toStruct("Object",JSON.parse(file));
console.log("finish")
  
