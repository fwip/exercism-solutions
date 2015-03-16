var Beer = function(){ };

function bottles(number){
  if (number > 1) {
    return String(number) + " bottles";
  } else if (number === 1) {
    return "1 bottle";
  }
  return "No more bottles";
}

function joiner(number){
  if (number > 1){
    return "Take one down and pass it around, ";
  } else if (number === 1) {
    return "Take it down and pass it around, ";
  } else {
    return "Go to the store and buy some more, ";
  }
}

Beer.prototype.verse = function(number){
  var next = number - 1;
  if (next < 0) {
    next = 99;
  }
  return bottles(number) + " of beer on the wall, " +
         bottles(number).toLowerCase() + " of beer.\n" +
         joiner(number) +
         bottles(next).toLowerCase() + " of beer on the wall.\n";
};

Beer.prototype.sing = function(start, end){
  end = end || 0; // If falsey, use zero to finish the song
  var verses = [];
  for (var i = start; i>=end; i--){
    verses.push(this.verse(i));
  }
  return verses.join("\n");
};

module.exports = new Beer();
