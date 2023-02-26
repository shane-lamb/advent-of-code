file_path = File.join(__dir__, "../problems", "day10", "input.txt")
file = File.open(file_path, "r")
instructions = file.each_line.map do |line|
  command, value = line.split " "
  next command == "addx" ? value.to_i : 0
end

current_value = 1
pending_value = 0
cycles = instructions.flat_map do |add_amount|
  current_value += pending_value
  pending_value = add_amount
  next [current_value] if add_amount == 0
  next [current_value, current_value]
end

# puts cycles.map.with_index { |value, index| "#{index + 1}. #{value}" }

# problem 1
sum = cycles.map.with_index do |value, index|
  cycle = index + 1
  on_interval = (cycle - 20) % 40 == 0
  if on_interval and cycle < 221
    puts "cycle #{cycle} value #{value}"
    next value * cycle
  end
  next 0
end.sum
puts "sum: #{sum}"

# problem 2
crt_width = 40
cycles.each_slice(crt_width).each do |cycles_batch|
  line = cycles_batch.map.with_index do |middle_sprite_pos, column|
    diff = (middle_sprite_pos - column).abs
    next diff > 1 ? "." : "#"
  end
  puts line.join
end
