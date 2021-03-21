def check_number(n)
  case n
  when n.odd?
    'Weird'
  when n.even? && n >= 2 && n <= 5
    'Not Weird'
  when n.even? && n >= 6 && n <= 20
    'Weird'
  else
    'Not Weird'
  end
end

