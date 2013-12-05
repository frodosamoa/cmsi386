// queue.js

function Queue() {
    var data = [];
    
    return {
        add: function (x) {data.push(x);},
        remove: function () {return data.shift();}
    };
}

module.exports = Queue;

// In the Node.js REPL...

> var queue = require('./queue.js')
undefined
> var q = queue()
undefined
> q
{ add: [Function], remove: [Function] }
> q.add(7)
undefined
> q.data
undefined
> q.remove()
7