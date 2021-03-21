defmodule Counter do
  def default(input) do
    [number_of_times | list] = input

    Enum.each(list, fn(item) ->
      Enum.each(1..number_of_times, fn _ ->
        IO.puts item
      end)
    end)
  end

  def recursive(_, []), do: 0
 
  def recursive(n, [head | tail]) do
    for _ <- 0..n, do: IO.puts head
    recursive(n, tail)
  end
end

defmodule Solution do
  def main() do
    #input = IO.read(:stdio, :all)
    #        |> String.split
    #        |> Enum.map(&String.to_integer(&1))
    #
    [head | tail] = [3, 1, 2, 3]
    
    #input =  [3, 1, 2, 3]
    Counter.default(input)
    #Counter.recursive(head - 1, tail)
  end
end

Solution.main
