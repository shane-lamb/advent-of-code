class Position
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def to_s
    "#{@x},#{@y}"
  end

  def hash
    [@x, @y].hash
  end

  alias eql? ==

  def ==(other)
    @x == other.x && @y == other.y
  end

  def +(other)
    Position.new(@x + other.x, @y + other.y)
  end

  def -(other)
    Position.new(@x - other.x, @y - other.y)
  end

  def abs
    Position.new(@x.abs, @y.abs)
  end
end

