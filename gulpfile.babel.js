var gulp   = require("gulp");
var gulpgo = require("gulp-go");

var go;

gulp.task("go-run", () => {
  go = gulpgo.run("./main/main.go", ["--arg1", "value1"], {cwd: __dirname, stdio: 'inherit'});
});

gulp.task("devs", ["go-run"], () =>  {
  gulp.watch([__dirname+"/**/*.go"]).on("change", function() {
    go.restart();
  });
});

gulp.task('default', ['devs'])