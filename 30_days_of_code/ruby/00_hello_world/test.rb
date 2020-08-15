require_relative 'main'

describe '#print_text' do
  it 'returns an array of phrases' do
    expected = ['Hello, Word', 'Welcome to 30 Days of Code!']
    result = print_text('Hello, Word', 'Welcome to 30 Days of Code!')

    expect(result).to eq(expected)
  end
end
