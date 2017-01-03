package main

import (
	_ "bufio"
	"fmt"
	"math"
	_ "os"
)

func main() {

	var S, K, T, R, V, d1, d2 float64
	var c, p map[string]float64
	greeks := make(map[string]float64)

	fmt.Scanln(&S, &K, &T, &R, &V)
	fmt.Println(greeks)
	fmt.Println(d1, d2, c, p)
	calculateGreeks(35, 40, 0.25, 0.03, 0.4, 0, 0, greeks)
	fmt.Println(greeks)
}

func calculateGreeks(S, K, T, R, V, d1, d2 float64, greeks map[string]float64) {

	d1 = (math.Log((S)/K) + (R+V*V/2.)*T) / (V * math.Sqrt(T))
	fmt.Println(d1)

	d2 = d1 - V*math.Sqrt(T)
	fmt.Println(d2)

	dis := gaussian.NewGaussian(0, 1)

	greeks["callPrice"] = S*dis.Cdf(d1) - K*math.Exp(-R*T)*dis.Cdf(d2)
	greeks["putPrice"] = K*math.Exp(-R*T) - S + greeks["callPrice"]

	greeks["callDelta"] = dis.Cdf(d1)
	greeks["callGamma"] = (dis.Pdf(d1)) / (S * V * math.Sqrt(T))
	greeks["callTheta"] = (-(S*V*dis.Pdf(d1))/(2*math.Sqrt(T)) - R*K*math.Exp(-R*T)*dis.Cdf(d2)) / 365
	greeks["callVega"] = S * math.Sqrt(T) * dis.Pdf(d1) / 100
	greeks["callRho"] = K * T * math.Exp(-R*T) * dis.Cdf(d2) / 100

	greeks["putDelta"] = (greeks["callDelta"]) - 1
	greeks["putGamma"] = greeks["callGamma"]
	greeks["putTheta"] = (-(S*V*dis.Pdf(d1))/(2*math.Sqrt(T)) + R*K*math.Exp(-R*T)*dis.Cdf(-d2)) / 365
	greeks["putVega"] = greeks["callVega"]
	greeks["putRho"] = -K * T * math.Exp(-R*T) * dis.Cdf(-d2) / 100

	return
}