$head_loc = [0, 0]
$tail_loc = [0, 0]
$visited = {}

def direction(number)
  number == 0 ? 0 : number / number.abs
end

def head_moves(direction)
  case direction
  when "R"
    $head_loc[0] += 1
    puts "head moves right"
  when "L"
    $head_loc[0] -= 1
    puts "head moves left"
  when "U"
    $head_loc[1] += 1
    puts "head moves up"
  when "D"
    $head_loc[1] -= 1
    puts "head moves down"
  end

  diff = $head_loc.zip($tail_loc).map { |a, b| a - b }
  puts "diff: #{diff}"

  if diff[0].abs > 1 || diff[1].abs > 1
    $tail_loc[0] += direction(diff[0])
    $tail_loc[1] += direction(diff[1])
  end

  $visited[$tail_loc.to_s] = true
end

file_path = File.join(__dir__, "../problems", "day9", "input.txt")
file = File.open(file_path, "r")
file.each_line do |line|
  direction, count = line.split " "
  count.to_i.times { head_moves(direction) }
end

puts $visited.length
