package physics

import vec "kepler/vec"

type kinematic struct {
	position vec.Vec
	velocity vec.Vec
}

type Particle struct {
	movement kinematic
	mass float64
	radia float64
}

func dualCollision(p1 Particle, p2 Particle, timeStep float64) (Particle , Particle) {
	// börjar med att kolla om de kolliderar
	// hitta den nädmaste punkten och kolla om det är mer eller mindre än
	// den kombinerade radien av de båda partiklarna
	// kan jag göra det så enkel??
	// får jag tekniska problem om de korsar varandra i tidsspannet?

	// om man kollar på vektorn mellan de två vektorerna så kan man använda intervall halvering
	// för att hitta punkten som de är närmast varandra,
	// notera att vektorn mellan dem roterar samtidigt som partiklarna rör sig mot eller ifrån varandra
	//

	//
}
