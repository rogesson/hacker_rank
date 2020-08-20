require_relative('main')

describe 'Person' do
  describe '#age' do
    it 'returns zero if age is less than zero' do
      person = Person.new(-1)
      expect(person.age).to eq(0)
    end
  end

  describe '#am_old?' do
    it 'returns young message if age is < 13' do
      person = Person.new(12)

      expect(person.am_old?).to eq('You are young.')
    end

    it 'returns teenager message if age is >= 13' do
      person = Person.new(14)

      expect(person.am_old?).to eq('You are a teenager.')
    end

    it 'returns old message if age is >= 18' do
      person = Person.new(18)

      expect(person.am_old?).to eq('You are old.')
    end
  end

  describe '#year_passes' do
    it 'increments age by 1' do
      person = Person.new(22)

      expect(person.age).to eq(22)

      person.year_passes

      expect(person.age).to eq(23)
    end
  end
end
