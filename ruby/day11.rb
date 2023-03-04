file_path = File.join(__dir__, "../problems", "day11", "input.txt")
file = File.open(file_path, "r")

Monkey = Struct.new("Monkey", :items, :operator, :operatee, :divisible_by, :true_monkey, :false_monkey, :inspections)

monkeys = file.each_line.reduce([[]]) do |monkey_lines, line|
  last = line == "\n" ? [] : monkey_lines.pop << line
  next monkey_lines << last
end.map do |lines|
  # Monkey 0:
  lines.shift
  # Starting items: 84, 66, 62, 69, 88, 91, 91
  starting_items = lines.shift.scan(/[0-9]+(?=,|$)/).map(&:to_i)
  # Operation: new = old * 11
  operator, operatee = lines.shift.scan(/(?<=\= old ).*/).first.split(" ")
  # Test: divisible by 2
  divisible_by = lines.shift.split(" ").last.to_i
  #   If true: throw to monkey 4
  true_monkey = lines.shift.split(" ").last.to_i
  #   If false: throw to monkey 7
  false_monkey = lines.shift.split(" ").last.to_i
  next Monkey.new(starting_items, operator, operatee, divisible_by, true_monkey, false_monkey, 0)
end

def do_turn(monkey, monkeys, worry_level_reduction_factor, regulation_value)
  while monkey.items.any?
    item = monkey.items.shift
    operatee = monkey.operatee == "old" ? item : monkey.operatee.to_i
    worry_level = monkey.operator == "*" ? item * operatee : item + operatee
    monkey.inspections += 1
    worry_level = worry_level / worry_level_reduction_factor
    worry_level = worry_level % regulation_value
    is_divisible = worry_level % monkey.divisible_by == 0
    target = is_divisible ? monkey.true_monkey : monkey.false_monkey
    monkeys[target].items << worry_level
  end
end


worry_level_reduction_factor = 1 # should = 3 for exercise 1
iterations = 10_000 # should = 20 for exercise 1

regulation_value = monkeys.map(&:divisible_by).reduce(:*)
puts monkeys
iterations.times do |i|
  monkeys.each { |monkey| do_turn(monkey, monkeys, worry_level_reduction_factor, regulation_value) }
  puts "After round #{i + 1}:"
  puts monkeys
end
puts monkeys.dup.sort_by { |monkey| -monkey.inspections}.take(2).map(&:inspections).reduce(:*)
