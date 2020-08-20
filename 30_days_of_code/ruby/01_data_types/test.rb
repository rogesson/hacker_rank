require_relative 'main'

describe '#sum' do
  it 'returns de sum of variables' do
    expected = [
      16,
      8.0,
      'HackerRank is the best place to learn and practice coding!'
    ]

    i = 4
    d = 4.0
    s = 'HackerRank '

    result = sum(i, d, s)

    expect(result).to eq(expected)
  end
end
