const print = @import("std").debug.print;
const fs = @import("std").fs;
const io = @import("std").io;
const mem = @import("std").mem;
const allocator = @import("std").heap.page_allocator;
const arrayList = @import("std").ArrayList(i32);
const fmt = @import("std").fmt;
const sort = @import("std").sort;
const autoHashMap = @import("std").AutoHashMap(i32, i32);

pub fn main() !void {
    //const file = try fs.cwd().openFile("day1_input_small", .{});
    const file = try fs.cwd().openFile("day1_input", .{});
    defer file.close();

    var bufferReader = io.bufferedReader(file.reader());
    var reader = bufferReader.reader();

    var left = arrayList.init(allocator);
    defer left.deinit();
    var right = arrayList.init(allocator);
    defer right.deinit();

    var buff: [1024]u8 = undefined;
    while (try reader.readUntilDelimiterOrEof(&buff, '\n')) |line| {
        var it = mem.splitSequence(u8, line, "   ");
        try left.append(try fmt.parseInt(i32, it.next().?, 10));
        try right.append(try fmt.parseInt(i32, it.next().?, 10));
    }

    var frequencyMap = autoHashMap.init(allocator);
    defer frequencyMap.deinit();

    for (right.items) |r| {
        try frequencyMap.put(r, (frequencyMap.get(r) orelse 0) + 1);
    }

    var total: i32 = 0;
    for (left.items) |l| {
        total += l * (frequencyMap.get(l) orelse 0);
    }
    print("{d}\n", .{total});
}
