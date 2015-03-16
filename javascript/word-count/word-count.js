var words = function(phrase){
  var counts = {};
  var words = phrase.split(/\s+/);

  words.forEach(function(word){
    if (counts.hasOwnProperty(word)){
      counts[word]++;
    } else {
      counts[word] = 1;
    }
  });
  return counts;
};

module.exports = words;
