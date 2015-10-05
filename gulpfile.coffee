gulp = require 'gulp'
browserSync = require 'browser-sync'
stylus = require 'gulp-stylus'
coffee = require 'gulp-coffee'
plumber = require 'gulp-plumber'
uglify = require 'gulp-uglify'
rename = require 'gulp-rename'
nib = require 'nib'
del = require 'del'
{argv} = require 'yargs'

project =
  name: 'Heron'
  src: 'assets'
  dest: 'static'
styles =
  name: 'styles'
  exts: ['styl', 'css']
scripts =
  name: 'scripts'
  exts: ['coffee', 'js']
templates =
  files: ->
    "templates/**/*.html"
assets =
  name: 'assets'
  dirs: [styles.name, scripts.name]
  exts: [].concat(styles.exts, scripts.exts)
  glob: ->
    dirs = assets.dirs.join(',')
    exts = assets.exts.join(',')
    "#{project.src}/{#{dirs}}/**/*.{#{exts}}"

gulp.task 'clean', ['clean:dist']
gulp.task 'build', ['stylus', 'coffee', 'collect']
gulp.task 'default', ['clean'], ->
  gulp.start 'build'

gulp.task 'watch', ['default'], ->
  gulp.start 'browser-sync'
  gulp.watch assets.glob(), ['build']
  gulp.watch templates.files(), ['reload']

gulp.task 'reload', ->
  gulp.src "#{templates.files()}"
    .pipe browserSync.reload
      stream: true  

gulp.task 'stylus', ->
  options =
    use: nib()
    compress: not argv.debug
  gulp.src "#{project.src}/**/*.styl"
    .pipe stylus options
    .pipe gulp.dest "#{project.dest}"
    .pipe browserSync.reload
      stream: true

gulp.task 'coffee', ->
  options =
    bare: true
  stream = gulp.src "#{project.src}/**/*.coffee"
    .pipe plumber()
    .pipe coffee options
  unless argv.debug
    stream = stream.pipe uglify()
  stream.pipe gulp.dest "#{project.dest}"
    .pipe browserSync.reload
      stream: true

gulp.task 'collect', ->
  gulp.src "#{project.src}/**/*.{js,css}"
    .pipe gulp.dest "#{project.dest}"

gulp.task 'browser-sync', ->
  browserSync
    port: argv.port + 100
    proxy: "127.0.0.1:#{argv.port}"
    open: false

gulp.task 'clean:dist', (done) ->
  del [
    "#{project.dest}/**/*.{#{assets.exts.join(',')}}"
  ], done
