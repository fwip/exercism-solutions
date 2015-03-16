module.exports = {
  verse: function(num){

    var meal = this.animals[num-1];
    var v = "I know an old lady who swallowed a " + meal.name + ".\n";
    if ("rhyme" in meal){
      v += meal.rhyme + "\n";
    }

    // If we're at horse, end early! (It makes the song much funnier.)
    if (num == this.animals.length){
      return v;
    }

    // Should yank out to its own function?
    for (var i=num-1; i>0; i--){
      var thisMeal = this.animals[i];
      var prevMeal = this.animals[i-1];
      v += "She swallowed the " + thisMeal.name + " to catch the " + prevMeal.name;

      // Spider has some extra wriggly going on there
      if ("extra" in prevMeal){
        v += prevMeal.extra;
      }

      v += ".\n";
    }

    v += "I don't know why she swallowed the fly. Perhaps she'll die.\n";
    return v;
  },

  verses: function(start, end){
    var v = "";
    // It'd be cool to write this as a map().join() expression
    for (var i=start; i<=end; i++){
      v += this.verse(i) + "\n";
    }
    return v;
  },

  animals: [
    {name: "fly"},
    {name: "spider", rhyme: "It wriggled and jiggled and tickled inside her.", extra: " that wriggled and jiggled and tickled inside her"},
    {name: "bird",   rhyme: "How absurd to swallow a bird!"},
    {name: "cat",    rhyme: "Imagine that, to swallow a cat!"},
    {name: "dog",    rhyme: "What a hog, to swallow a dog!"},
    {name: "goat",   rhyme: "Just opened her throat and swallowed a goat!"},
    {name: "cow",    rhyme: "I don't know how she swallowed a cow!"},
    {name: "horse",  rhyme: "She's dead, of course!"},
  ],
};
