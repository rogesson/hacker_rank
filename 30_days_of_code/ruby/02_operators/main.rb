def total_cost(meal_cost, tip_percent, tax_percent)
  tip = meal_cost * tip_percent / 100.0
  tax = meal_cost * tax_percent / 100.0

  (meal_cost + tip + tax).round
end

