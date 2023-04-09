~/bin/ffmpeg -framerate 5 -pattern_type glob -i '*.jpg' -vf scale=-2:1080,format=yuv420p -movflags +faststart output.mp4
