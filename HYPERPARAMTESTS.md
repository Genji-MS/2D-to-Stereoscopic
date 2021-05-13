# Argument.yml Paramater testing
Given the notes from our Tests, we want to adjust some of the settings to get better results from our images.


## Depth Map
It became obvious after a couple of adjustments that the settings in argument.yml are only for the rendering of our image. The depth maps created by MiDaS are identical to previous attempts. Therefore, images where detail is not generated at the appropriate depth cannot be fixed by these parameters.

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Depth.png" height="350">


## Parameter testing
For these tests I selected the one image that would have benefited the most, the hydrant. Because the flowers in the foreground were flat. At the time, I had hoped to bring those details out. Instead, it just became a study of what these parameters do.

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/class/Term9/ML/2D-to-Stereoscopic/Tests/StereoOutput/2D_Test_Hydrant_CrossView.png" height="350">

### `depth_threshold` Original: 0.04
    - A threshold in disparity, adjacent two pixels are discontinuity pixels if the difference between them excceed this number.

    Smaller: 0.01                            Greater: 0.1
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_DT_0-01_CrossView.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_DT_0-1_CrossView.png" height="350">

    Notes:
- A wider range of settings were tested with no significant differences. 0.005 did not render
- Higher values were almost a minute faster to render
- 

### `ext_edge_threshold` Original: 0.002
    - The threshold to define inpainted depth edge. A pixel in inpainted edge map belongs to extended depth edge if the value of that pixel exceeds this number,

    Smaller: 0.001                            Greater: 0.1
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_EET_0-01_CrossView.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_EET_0-1_CrossView.png" height="350">

    Notes:
- At higher values, the tops of trees are trimmed
- 
    
### `redundant_number` Original: 12
    - The number defines short segments. If a depth edge is shorter than this number, it is a short segment and removed.

    Smaller: 6                               Greater: 24
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_RN_6_CrossView.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_RN_24_CrossView.png" height="350">

    Notes:
- ?

### `background & context thickness` Original 70, 140, 70, 70
    - The thickness of ... area. (first and second inpaint process)

    Smaller: 35,70,35,35                     Greater: 140,280,140,140
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Thick_Half_CrossView.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/Params/2D_Test_Hydrant_Thick_Dbl_CrossView.png" height="350">

    Notes:
- Greater took 3:56 to process
- Shorter took 2:50 to process


## Left Input
A selection of images were thought they would look better if the left side of the image pair were used. For comparison the previously generated image is also shown.

###  -   Mouth
    Left Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Mouth.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Mouth.png" height="350">  

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Mouth_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mouth_CrossView.png" height="350">

###  -   Overview
    Left Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Overview.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Overview.png" height="350">  

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Overview_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Overview_CrossView.png" height="350">

###  -   Plant

    Left Image
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Plant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Plant.png" height="350">  

    Generated from Left
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/LeftInput/2D_Test_Plant_CrossView.png" height="350"> 

    Generated from Right (Previous)
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Plant_CrossView.png" height="350">


## Disparity
These images were generated with a greater disparity (0.0175) In an attempt to either recreate an image closer to the original, or to see if it would simply look better

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
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Art_Bart.jpg" height="350">

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Art_Bart_CrossView.png" height="350">

    Increased Disparity
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/ParameterTuned/PlusDisparity/2D_Art_Bart_CrossView_D.png" height="350">