input_string = gets

def print_text(*args)
  args.each do |text|
    puts text
  end
end

print_text('Hello, World', input_string)
