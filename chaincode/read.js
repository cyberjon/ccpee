router.post('/sendTransaction', function(req, res, next) {
    console.log("savedata called: " + req.body + " ----------saveTransaction-------------- ");
    


    
    var count = 0;
    for (key in req.body)              // should return 2
    {
       if(req.body.hasOwnProperty(key))
       {
          count++;
       }
    }

    var last = req.body.length;

    var result_data = [];
    var k = 0;
    for(var key in req.body) {
        console.log("key: ");
        console.log(key);
        console.log("k: ");
        console.log(k);
        if(req.body.hasOwnProperty(key)) {
            transaction = req.body[key];
            console.log("Transactions received: ");
            console.log(transaction);
            console.log("key after if: ", key);

            // Set values of json transaction
            var id = transaction.Order_id;
            var userA = transaction.User_A;
            var userB = transaction.User_B;
            var seller = transaction.Seller; 
            var amount = transaction.Ex_points;
            var prev_trans_id = transaction.Prev_Transaction_ID;
            var exp_date = transaction.Exp_date;

            var exp_date = new Date(exp_date);
            var exp_dateStr = exp_date.getFullYear()+''+(exp_date.getMonth()+1)+''+exp_date.getDate();

            var curret_date = new Date();
            var dateStr = curret_date.getFullYear()+''+(curret_date.getMonth()+1)+''+curret_date.getDate();

            console.log("Console all data: ###################### ");
            console.log("##########################################");
            console.log("id: " + id + " userA: " + userA + " userB: " + userB + " seller: " + seller + " amount: " + amount + " prev_trans_id: " + prev_trans_id + " dateStr: " + dateStr + " exp_dateStr: " + exp_dateStr)
            console.log("##########################################");

            my_cc.invoke.init_transaction([id,userA,userB,seller,amount,prev_trans_id,dateStr, exp_dateStr],function(err, data) {
                console.log('Returned data success', data);
                var succ_data = data;
                //data = JSON.stringify(data);
                console.log("Inside the invoke key is: ", key);
                console.log("Inside the invoke k is: ", k);
                console.log('Lets push the json data into array -----------------------');
                console.log(result_data);
                result_data.push(succ_data);
                //result_data = JSON.stringify(result_data);

                if (result_data.length == last) {
                    console.log("###### Key = last ####### ");                    
                    console.log("key", key);
                    console.log("last", last);
                    console.log("k", k);
                    console.log("###########################")
                    result_data = JSON.stringify(result_data);
                    res.json(result_data);
                }  
            });
        }
        
    }
    
});