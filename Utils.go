package utils

import (
	"sync"
)

type Pair[T any, U any] struct {
	first  T
	second U
}

func ForSquare(size int, action func(int, int)) {
	For(size, func(i int) {
		For(size, func(j int) {
			action(i, j)
		})
	})
}

func For(size int, action func(int)) {
	for i := 0; i < size; i++ {
		action(i)
	}
}

func ForEach[T any](arr []T, action func(T)) {
	for _, a := range arr {
		action(a)
	}
}

func Map[T any, U any](arr []T, action func(T) U) (output []U) {
	for _, a := range arr {
		output = append(output, action(a))
	}
	return
}

func Filter[T any](arr []T, test func(T) bool) (output []T) {
	for _, a := range arr {
		if test(a) {
			output = append(output, a)
		}
	}
	return
}

func ForEachWG[T any](arr []T, action func(T)) {
	var wg sync.WaitGroup
	poolSize := len(arr)
	wg.Add(poolSize)
	inputs := make(chan T, poolSize)
	for range arr {
		go func() {
			defer wg.Done()
			for input := range inputs {
				action(input)
			}
		}()
	}

	for _, a := range arr {
		inputs <- a
	}
	close(inputs)
	wg.Wait()
}

func MapWG[T any, U any](arr []T, action func(T) U) (result []U) {
	var wg sync.WaitGroup
	poolSize := len(arr)
	wg.Add(poolSize)
	inputs := make(chan T, poolSize)
	outputs := make(chan U, poolSize)

	for range arr {
		go func() {
			defer wg.Done()
			for input := range inputs {
				outputs <- action(input)
			}
		}()
	}

	for _, a := range arr {
		inputs <- a
	}
	close(inputs)

	go func() {
		wg.Wait()
		close(outputs)
	}()

	for o := range outputs {
		result = append(result, o)
	}
	return result
}

func OrderedMapWG[T any, U any](arr []T, action func(T) U) []U {
	var wg sync.WaitGroup
	poolSize := len(arr)
	wg.Add(poolSize)
	inputs := make(chan Pair[int, T], poolSize)
	outputs := make(chan Pair[int, U], poolSize)
	result := make([]U, len(arr))

	for range arr {
		go func() {
			defer wg.Done()
			for input := range inputs {
				outputs <- Pair[int, U]{
					first:  input.first,
					second: action(input.second),
				}
			}
		}()
	}

	for i, a := range arr {
		inputs <- Pair[int, T]{
			first:  i,
			second: a,
		}
	}
	close(inputs)

	go func() {
		wg.Wait()
		close(outputs)
	}()

	for o := range outputs {
		result[o.first] = o.second
	}
	return result
}

func FilterWG[T any](arr []T, test func(T) bool) (result []T) {
	var wg sync.WaitGroup
	poolSize := len(arr)
	wg.Add(poolSize)
	inputs := make(chan T, poolSize)
	outputs := make(chan T, poolSize)

	for range arr {
		go func() {
			defer wg.Done()
			for input := range inputs {
				if test(input) {
					outputs <- input
				}
			}
		}()
	}

	for _, a := range arr {
		inputs <- a
	}
	close(inputs)

	go func() {
		wg.Wait()
		close(outputs)
	}()

	for o := range outputs {
		result = append(result, o)
	}
	return result
}

func ActionWG(actions []func()) {
	var wg sync.WaitGroup
	poolSize := len(actions)
	wg.Add(poolSize)

	inputs := make(chan func(), poolSize)
	for range actions {
		go func() {
			defer wg.Done()
			for input := range inputs {
				input()
			}
		}()
	}

	for _, a := range actions {
		inputs <- a
	}
	close(inputs)
	wg.Wait()
}
