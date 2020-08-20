class Person
  attr_accessor :age

  def initialize(initial_age)
    if initial_age < 0
      initial_age = 0
      puts 'Age is not valid, setting age to 0.'
    end

    self.age = initial_age
  end

  def am_old?
    message = if age < 13
                'You are young.'
              elsif age >= 13 && age < 18
                'You are a teenager.'
              else
                'You are old.'
              end

    puts message
    message
  end

  def year_passes
    self.age += 1
  end
end
