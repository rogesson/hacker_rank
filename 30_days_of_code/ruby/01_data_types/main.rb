def sum(integer, decimal, string)
  arr = []
  arr << (integer * decimal).to_i
  arr << decimal * 2
  arr << string + 'is the best place to learn and practice coding!'

  arr.each do |a|
    puts a
  end
end

i = 4
d = 4.0
s = 'HackerRank '

sum(i, d, s)
