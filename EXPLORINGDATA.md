# Exploring the model and Data

Unlike most notebooks, this model is almost entirely excuted in python scripts that are imported into Google Colab. Our models are pretrained, so instead of looking at the individual datasets used to train these models I will explore what these individual models do for our process and how to navigate the repo. 

The models are referenced in the argument.yml file. This file has no comments, however DOCUMENTATION.md fills in the necessary details. https://github.com/vt-vl-lab/3d-photo-inpainting/blob/master/DOCUMENTATION.md

- `MiDaS_model_ckpt: MiDaS/model.pt`

  "MiDaS was trained on 10 datasets (ReDWeb, DIML, Movies, MegaDepth, WSVD, TartanAir, HRWSI, ApolloScape, BlendedMVS, IRS) with multi-objective optimization." 
  - It's objective is to create a depth map of a 2D image. It's a grey scale image where white pixels are closer to the camera, and black pixels further. We use this grey scale image to create a 3D wireframe of our 2D image.

- `depth_edge_model_ckpt: checkpoints/EdgeModel.pth`
- `rgb_feat_model_ckpt: checkpoints/ColorModel.pth`

  "image inpainting recreates how artists work: lines first, color next." 
  - Together, these models create the textured 3D object

- `depth_feat_model_ckpt: checkpoints/DepthModel.pth`

  "This is the model being demonstrated in the referenced paper above. It leverages the other models to create our final 3D output"
  - This model fills in 'missing' image data for a 3D scene. There are obvious holes when we move our camera in 3D space. The technique here fills in these holes better than competators.


## Modifications
#### The following changes are made in `argument.yml` to create Stereoscopic images, instead of video output

- `num_frames: 2`
    - We only need to process 2 images to create a stereoscopic image
- `x_shift_range: [0.0075,-0.0075]`
    - Disparity between the left and right images on the X-axis. 
      - Large values EX 0.1 will create horrible images!
      - A range around 0.01 works great. 0.075 creates less disparity and feels comfortable for my own eyes.
        - Transforming to the right (positive value) creates a 'CrossView' type of image. 
        - Transforming to the left (negative value) creates a 'ParallelView' type of image.
- `y_shift_range: [0.00,0.00]`
    - The original paper uses these to create a spiral video effect on the Y-axis. Unused for Stereoscopic images.
- `z_shift_range: [0.00,0.00]`
    - The original paper uses these to create a spiral video effect on the Z-axis. Unused for Stereoscopic images.
- `traj_types: ['double-straight-line', 'double-straight-line']`
    - The type of camera trajectory. `straight-line` or `circle`
- `video_postfix: ['Cross', Parallel']`
    - These entries correspond to video effects. For our purpose, they append to the base image name when saving the image. 
- Note that the number of elements in `x_shift_range`,  `y_shift_range`, `z_shift_range`, `traj_types` and `video_postfix` should be equal.
    - Adding additional items to these lists would allow our model to process additional effects, and additional output files for each 2D image. 

#### The following changes have been inserted into `mesh.py` to process our images into a single stereoscopic image and removes the creation of an MP4 video.

from PIL import Image
  - Using the Pillow library to process images
...
clip.write_videofile(os.path.join(output_dir, video_basename + '_' + video_traj_type + '.mp4'), fps=config['fps'])
  - The next to last line of the script, we comment this out as we don't want to process mp4 videos.
path = os.path.join(output_dir, video_basename + '_' + video_traj_type + '.png')
  - store the path for our final image
img_left_arr, img_right_arr = stereos[0],stereos[1]
  - `stereos` defined by "num_frames", an array holding images for processing into a video.
y_size, x_size, _ = img_left_arr.shape
  - Store the Width and Height of our image, the 3rd param is color channels. NOTE: Height is before Width.
stereo_file = Image.new("RGB", (x_size*2,y_size))
  - Create a new image with double the width to hold both the left/right stereoscopic images
img_left = Image.fromarray(img_left_arr)
img_right = Image.fromarray(img_right_arr)
  - Data is stored in a numpy array. Not to be mistaken as image data, we call on `Pillow.Image` to read it.
stereo_file.paste(img_left)
stereo_file.paste(img_right,(x_size+1,0))
  - Copies our image data into the left and right sides of our stereoscopic image
stereo_file.save(path)
  - Saves the file

## Visualizing our Data

Input Image
  - The file is placed in the /image/ folder to be processed.
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/image/Flower.jpg" height="250">

Depth Map
  - The first step is using MiDaS to generate our depth map
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/depth/Flower.png" height="250">

3D Model from Depth Map
  - The depth map is used to create a 3D object, mapping the image as a texture over the 3D object. Note, the "holes" that are visible behind the objects, these are what get filled in by the model in the referenced paper.
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/depth/Model1.png" height="250">

Output Image 
  - By shifting our object in 3D space, filling the holes and generating two images, we've created a CrossView Stereoscopic image.
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/video/Flower_CrossView.png" height="250">

Gif
  - Showing the disparity between the two images. 
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/video/Flower.gif">


# Visualising Stereoscopic images

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/depth/StereoTypes.png" height = "200">

Types of Stereoscopic Images

- CrossView. You look at a point in FRONT of the image. Your left eye looks at the right image.
- ParallelView. You look at a point BEYOND the image. Your left eye looks at the left image.

Viewing Stereoscopic images with your eyes goes beyond the scope of this repo. But here are some resources:
https://www.reddit.com/r/CrossView/

https://www.reddit.com/r/ParallelView/