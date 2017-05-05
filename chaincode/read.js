router.post('/getData', function(req, res, next){
    var id = req.body.trId;
    my_cc.query.read(['read',id],function(err,resp){
        if(!err){
            var pre = JSON.stringify(resp);
            console.log('success',pre);  
			return pre;
        }else{
            console.log('fail');
        }
    });
});