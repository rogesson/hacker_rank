require_relative('main')

describe '#check_number' do
  it 'returns weird when number is odd' do
    result = check_number(3)
    expect(result).to eq('Weird')
  end

  it 'returns Not weird when number is even and more than 20' do
    result = check_number(24)
    expect(result).to eq('Not Weird')
  end
end
