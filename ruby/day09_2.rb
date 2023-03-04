require_relative "lib/position"

$knots = Array.new(10) { Position.new(0, 0) }
$visited = {}

def direction(number)
  number == 0 ? 0 : number / number.abs
end

$offset_map = {
    "R" => Position.new(1, 0),
    "L" => Position.new(-1, 0),
    "U" => Position.new(0, 1),
    "D" => Position.new(0, -1)
}

def head_moves(direction)
  $knots[0] += $offset_map[direction]

  (1..$knots.length - 1).each do |index|
    diff = $knots[index - 1] - $knots[index]

    if diff.x.abs > 1 || diff.y.abs > 1
      $knots[index] += Position.new(direction(diff.x), direction(diff.y))
    end
  end

  $visited[$knots.last] = true
end

file_path = File.join(__dir__, "../problems", "day9", "input.txt")
file = File.open(file_path, "r")
file.each_line do |line|
  direction, count = line.split " "
  count.to_i.times { head_moves(direction) }
end

puts $visited.length
