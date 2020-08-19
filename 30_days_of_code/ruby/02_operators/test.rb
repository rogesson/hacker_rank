require_relative 'main'

describe '#total_cost' do
  it 'returns the total cost value' do
    result = total_cost(12.00, 20, 8)
    expect(result).to eq(15)
  end
end
