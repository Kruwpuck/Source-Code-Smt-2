// Buatlah prosedur saja

func jumlahPushSitUP(P,S,X,Y,B int, jumP, jumS *int){
    /*I.S. Terdefinisi bilangan bulat P,S,X,Y,B. jumlahPushUp dan jumlahSitUp diinisialiasi dengan bilangan 0.
    F.S. Menyimpan total push-up dan total sit-up pada variaebl jumlahPushUp dan jumlahSitUp
    */
    var i int
    

    for i = 1; i <= B*30; i++{
        if i % X == 0 || i % Y == 0{
            }else if i % 2 != 0{
                *jumP = *jumP + P
            } else if i % 2 == 0 {
                 *jumS = *jumS + S
            }  
        }
    }
