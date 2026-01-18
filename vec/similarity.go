package vec

func (v Vector) CosineSimilarity(u Vector) float64 {
	return v.DotProduct(u) / (v.Magnitude() * u.Magnitude())
}
