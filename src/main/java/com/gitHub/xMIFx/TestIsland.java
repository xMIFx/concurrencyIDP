package com.gitHub.xMIFx;

public class TestIsland {
	public static void main(String[] args) {
		byte[][] map = {
				{1, 0, 0, 1},
				{1, 0, 1, 1},
				{1, 0, 1, 1}
		};
		boolean[][] checkedMap = createCheckedMap(map);
		int count = 0;
		while (true) {
			int[] uncheckedIndexis = getUncheckedIndex(checkedMap);
			if (uncheckedIndexis == null) {
				break;
			}
			int firstInd = uncheckedIndexis[0];
			int secondInd = uncheckedIndexis[1];
			checkedMap[firstInd][secondInd] = true;
			//TODO change ind -1, ind +1
			if (firstInd > 0 && checkedMap[firstInd-1][secondInd]) {
				continue;
			} else if (secondInd > 0 && checkedMap[firstInd][secondInd-1]){
				continue;
			}
			System.out.println(uncheckedIndexis[0] + "_" + uncheckedIndexis[1]);
			count++;

		}
		System.out.println(count);
	}

	private static int[] getUncheckedIndex(boolean[][] checkedMap) {
		for (int i = 0; i < checkedMap.length; i++) {
			for (int j = 0; j < checkedMap[i].length; j++) {
				if (!checkedMap[i][j]) {
					return new int[] {i, j};
				}
			}
		}
		return null;
	}

	private static boolean[][] createCheckedMap(byte[][] map) {
		boolean[][] checkedMap = new boolean[map.length][map[0].length];
		for (int i = 0; i < map.length; i++) {
			for (int j = 0; j < map[i].length; j++) {
				boolean checked = false;
				if (map[i][j] != 1) {
					checked = true;
				}
				checkedMap[i][j] = checked;
			}
		}
		return checkedMap;
	}
}
