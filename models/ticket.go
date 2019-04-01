package booking




//Get place in train car

//Train ID, Train routing, Date(Dep/Ariv), Time(Dep/Ariv), Duration(Time), Free Places(quantitu, num)
// select place num, quantity of selecting place

//Rout
// api/v1/select_place/ GET  {return quantity & number of palaces}
// api/v1/ticket_order/ POST {
//                send: book place in a train}
//             receive: pas1{UserName, TrainNum}
//                      pas2{UserName, TrainNum}
//                      pas3{UserName, TrainNum}
//                      pas4{UserName, TrainNum}
// api/v1/bucket/tickets/ GET
//                           pas1{UserName, TrainNum, Price}
//...                        pas2{UserName, TrainNum, Price}
//...........................SUM price
//                           set in Redis Timer 15min
// api/v1/bucket/tickets/id DELETE return
//...                        pas2{UserName, TrainNum, Price}
//...........................delete in redis
//                           Redis Timer continue
//.api/v1/payment/tickets/id

