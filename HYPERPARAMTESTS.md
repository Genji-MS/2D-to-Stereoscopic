# Paramater Testing
Given the notes from our Tests, we want to adjust some of the settings to get better results from our images. These were performed by either modifying parameters in the Argument.yml file, by generating an image from different source material, or by adjusting the disparity between images.


## Depth Map
It became obvious after a couple of adjustments that the settings in argument.yml are only for the rendering of our image. The depth maps created by MiDaS are identical to previous attempts. Therefore, images where detail is not generated at the appropriate depth cannot be fixed by these parameters.

    Original Image & Generated Depth Map
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Hydrant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Depth.png" height="350">


## Parameter testing
For these tests I selected the one image that would have benefited the most, the hydrant. Because the flowers in the foreground were flat. At the time, I had hoped to bring those details out. Instead, they became a study of what these parameters do.

    Generated Image using default params
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Hydrant_CrossView.png" height="350">

### `depth_threshold` Original: 0.04
A threshold in disparity, adjacent two pixels are discontinuity pixels if the difference between them excceed this number.

    Smaller: 0.01                            
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_DT_0-01_CrossView.png" height="350">

    Greater: 0.1
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_DT_0-1_CrossView.png" height="350">

    Notes:
- A wider range of settings were tested with no significant differences. 0.005 did not render
- Higher values were almost a minute faster to render
- The tall tree in the background above the building at center is missing in the lower threshold value

### `ext_edge_threshold` Original: 0.002
The threshold to define inpainted depth edge. A pixel in inpainted edge map belongs to extended depth edge if the value of that pixel exceeds this number,

    Smaller: 0.001                            
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_EET_0-01_CrossView.png" height="350">

    Greater: 0.1
 <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_EET_0-1_CrossView.png" height="350">

    
### `redundant_number` Original: 12
The number defines short segments. If a depth edge is shorter than this number, it is a short segment and removed.

    Smaller: 6                               
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_RN_6_CrossView.png" height="350"> 

    Greater: 24
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_RN_24_CrossView.png" height="350">


### `background & context thickness` Original 70, 140, 70, 70
The thickness of ... area. (first and second inpaint process)

    Smaller: 35,70,35,35                    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Thick_Half_CrossView.png" height="350"> 

    Greater: 140,280,140,140
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Thick_Dbl_CrossView.png" height="350">

    Notes:
- Greater took 3:56 to process
- Shorter took 2:50 to process


## Left Input
A selection of images were thought they would look better if the left side of the image pair were used. For comparison the previously generated image is also shown.

###  -   Mouth

    DepthMaps Left & Right respectively
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Mouth.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Mouth.png" height="350">

    Original Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mouth.jpg" height="350">

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Mouth_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mouth_CrossView.png" height="350">


###  -   Overview

    DepthMaps Left & Right respectively
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Overview.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Overview.png" height="350">

    Original Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Overview.jpg" height="350">

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Overview_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Overview_CrossView.png" height="350">


###  -   Plant

    DepthMaps Left & Right respectively
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Plant.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Plant.png" height="350">

    Original Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Plant.jpg" height="350">

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Plant_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Plant_CrossView.png" height="350">


## Disparity
These images were generated with a greater disparity (0.0175) as an attempt to either recreate an image closer to the original, or to see if it would simply look better with a larger significance of depth

###  -   90596

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_90596.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_90596_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/2D_Test_90596_CrossView_D.png" height="350"> 


###  -   Mirror

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mirror.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mirror_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/2D_Test_Mirror_CrossView_D.png" height="350"> 


###  -   Mouth

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mouth.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mouth_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/2D_Test_Mouth_CrossView_D.png" height="350">


###  -   Swim

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Swim.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Swim_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/2D_Test_Swim_CrossView_D.png" height="350">


###  -   Bart

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Bart.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Bart_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/Art_Bart_CrossView_D.png" height="350">


## Conclusions and Observations

- The parameters in the argument.yml file appear to have small, almost imperceptable changes. However, I believe the problem to be poor image choice. An image with an object closer to the foreground would allow us to observe the details the inPaint model generates to fill in the gaps where there is no image data

- When generating a 3D image from a single image, some considerations need to be made. Like with the plant image, just a few pixels difference drastically changed the background building's texture. In terms of generating the depth map, MiDaS is exceptional, small changes in the photo do not appear to create any significant changes.

- Increasing Disparity gives more depth to the image, but futher considerations of cropping become apparent. The Swim image for example, the fingers are cropped out, making for a slightly unpleasant 3D experience. It comes down to a matter of personal preferance
