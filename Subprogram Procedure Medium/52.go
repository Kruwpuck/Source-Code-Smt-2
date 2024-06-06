func suhu(temp int) {
    /*I.S. terdefinisi sebuah bilangan bulat yang menyatakan temperatur.
    F.S. print string untuk menyatakan status cuaca berikut: "Freezing weather", 
    "Very Cold weather", "Cold weather", "Normal in Temp", "it's Hot", atau 
    "It's Very Hot".  sesuai dengan list cuaca diatas  
    */
    if temp < 0 {
        fmt.Print("Freezing weather") 
    }else if temp >=0 && temp <= 9 {
        fmt.Print("Very Cold weather")
    }else if temp >= 10 && temp <= 19 {
        fmt.Print("Cold weather")
    }else if temp >= 20 && temp <= 29 {
        fmt.Print("Normal in Temp") 
    }else if temp >= 30 && temp <= 39 {
        fmt.Print("It's Hot") 
    }else if temp >=40 {
        fmt.Print("It's Very Hot") 
}
}