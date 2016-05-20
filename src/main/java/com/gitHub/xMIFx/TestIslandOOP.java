package com.gitHub.xMIFx;

import java.util.ArrayList;
import java.util.List;

public class TestIslandOOP {
	private static byte[][] map = {
			{1, 0, 0, 1},
			{1, 0, 1, 1},
			{1, 1, 1, 1}
	};
	private static List<Position> islands = new ArrayList<>();

	public static void main(String[] args) {

		for (int i = 0; i < map.length; i++) {
			for (int j = 0; j < map[i].length; j++) {
				if (map[i][j] == 1) {
					Position position = new Position(i, j);
					if (islands.size() == 0) {
						islands.add(position);
					} else {
						boolean itOldIsland = false;
						for (Position p : islands) {
							itOldIsland = p.checkThatItOneIsland(position);
							if (itOldIsland) {
								break;
							}
						}
						if (!itOldIsland) {
							islands.add(position);
						}
					}
				}
			}
		}
		for (int i = 1; i < islands.size(); i++) {
			for (int j = 0; j < i; j++) {
				if (islands.get(j).checkThatItOneIsland(islands.get(i))) {
					islands.remove(i);
					i--;
				}
			}
		}

		System.out.println(islands);
		System.out.println(islands.size());
	}

	private static class Position {
		int y;
		int x;
		Position top;
		Position right;
		Position bottom;
		Position left;

		Position(int y, int x) {
			this.y = y;
			this.x = x;
		}

		boolean checkThatItOneIsland(Position position) {
			int deltaX = this.x - position.x;
			int deltaY = this.y - position.y;
			if (deltaY == -1 && deltaX == 0) {
				this.bottom = position;
				return true;
			}
			if (deltaY == 0 && deltaX == -1) {
				this.left = position;
				return true;
			}
			if (deltaY == 1 && deltaX == 0) {
				this.top = position;
				return true;
			}
			if (deltaY == 0 && deltaX == 1) {
				this.right = position;
				return true;
			}
			boolean old = checkAllCurrentPosition(position, false);
			old = checkAllChildPosition(position, old);
			return old;
		}

		private boolean checkAllChildPosition(Position position, boolean old) {
			if (!old && position.top != null) {
				old = this.checkThatItOneIsland(position.top);
			}
			if (!old && position.right != null) {
				old = this.checkThatItOneIsland(position.right);
			}
			if (!old && position.bottom != null) {
				old = this.checkThatItOneIsland(position.bottom);
			}
			if (!old && position.left != null) {
				old = this.checkThatItOneIsland(position.left);
			}
			return old;
		}

		private boolean checkAllCurrentPosition(Position position, boolean old) {
			if (top != null) {
				old = top.checkThatItOneIsland(position);
			}
			if (!old && right != null) {
				old = right.checkThatItOneIsland(position);
			}
			if (!old && bottom != null) {
				old = bottom.checkThatItOneIsland(position);
			}
			if (!old && left != null) {
				old = left.checkThatItOneIsland(position);
			}
			return old;
		}

		@Override
		public String toString() {
			return "Position{" +
					"x=" + x +
					", y=" + y +
					'}';
		}
	}
}
