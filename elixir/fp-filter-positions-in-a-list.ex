import Integer
defmodule Solution do
  def main() do
    input = IO.read(:stdio, :all)
            |> String.split
            |> Enum.map(&String.to_integer(&1))
    
    
    for index <- 0..length(input), is_odd(index), do: Enum.at(input, index) 
      |> IO.puts
  end
end

Solution.main
