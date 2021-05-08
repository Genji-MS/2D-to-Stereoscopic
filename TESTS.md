# Stereoscopic images vs single image to stereo generation
Test results are difficult to quantify numerically. Natural stereoscopic images don't have a specific disparity between images. Actually, hyperstereo images spanning thousands of miles apart exist on the internet. Recreating an identical disparity between images is not the goal of this model. So to compare our generated stereoimages with the originals by a numerical pixel perfect accuraccy would be inaccurate.

We compare the naturally created stereoscopic images against generated images by their visual inacuracies, noting any missing or innacurate details in the generated 3D scenes. 

The generated images use only the right side of the original set and use a `x_shift_range: 0.0075` for disparity. This value is global to all of our generated images. Images were cropped using the `ImageCrop` script included in this repo. Stereoscopic images were procurred from `reddit.com/r/CrossView/`

### Vs Test# 01  -   90596

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_90596.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_90596.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_90596_CrossView.png" height="350">

    Notes:
 - At first glance it appears her shirt/abdomen is too far forward. But comparing it with the original it is absolutely correct!
- The original image has too much disparity. 

### Vs Test# 02  -   Mirror

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mirror.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Mirror.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mirror_CrossView.png" height="350"> 

    Notes:
- The surface of the mirror lacks depth, there is some, but the interior phone has none
- The outer ring of the mirror doesn't have the same defraction. 
    - I do not think this is something a depth map can solve

### Vs Test# 03  -   Mouth

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mouth.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Mouth.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mouth_CrossView.png" height="350"> 

    Notes:
- The hair on top of his head has too much depth
- His mouth & nose have imperceptable depth.
    - This might look better with greater disparity 

### Vs Test# 04  -   Plant

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Plant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Plant.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Plant_CrossView.png" height="350"> 

    Notes:
- The three depths of plants are not generated correctly
- The buiding does not have the same texture
    - Since the image was generated from only the right image of the original pair, the original wall was not seen when generating the scene
    - This might generate better from the left image

### Vs Test# 05  -   Puddle

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Puddle.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Puddle.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Puddle_CrossView.png" height="350"> 

    Notes:
- The puddle does not generate the complete reflective depth
- The small objects in the exterior of the room did not generate depth
    - This might be fixed by modifying some internal settings

### Vs Test# 06  -   Xray

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Xray.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Xray.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Xray_CrossView.png" height="350"> 

    Notes:
- While there is some depth, the translucency of the skin is not generated

### Vs Test# 07  -   Hydrant

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Hydrant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Hydrant.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Hydrant_CrossView.png" height="350"> 

    Notes:
- The center flowers generate no depth
    - This might be fixed by modifying some internal settings


### Vs Test# 08  -   Dog     -   with cutouts

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Dog.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Dog.png" height="280"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Dog_CrossView.png" height="350"> 

    Notes:
- No depth of the smaller objects: leash, branch in mouth
- The popout effect from the image trim is not generated
    - This model was not designed for such effects
- The water does not generate depth

### Vs Test# 09  -   Overview     -   with cutouts

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Overview.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Overview.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Overview_CrossView.png" height="350"> 

    Notes:
- The popout effect from the image trim is not generated, 
    - This model was not designed for such effects

- There is some depth, but not as detailed
    - This might look better with greater disparity 
    - This might be fixed by modifying some internal settings for smaller details

### Vs Test# 10  -   Swim     -   with cutouts

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Swim.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Swim.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Swim_CrossView.png" height="350"> 

    Notes:
- The reflectivity/depth of the water is not generated

- The border going behind the hand made an attempt, but did not generate the correct level of depth.

- Not at much depth as the original
    - This might look better with greater disparity 

### Vs Test# 11  -   Text

    - Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Text.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Text.png" height="350"> 

    - Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Text_CrossView.png" height="350"> 

    Notes:
- The text pop-out effect is not generated, however it is artifically created by manipulating the disparity at specific areas of the original image pair
    - I would say this was generated correctly, I expected characters to generate at different depths, but they did not.

### Vs Test# 12  -   Art

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Art.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Art.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_CrossView.png" height="350"> 

    Notes:
- The objects in the background did not generate the same level of depth
    - This might look better with greater disparity, but I suspect it will not be 100% correct.

- The chandelier at top did not generate any depth

### Vs Test# 13  -   Meme

    Original
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Meme.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Meme.png" height="350"> 

    Generated
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Meme_CrossView.png" height="350"> 

    Top Comments:
- That look in his eyes made my day
    - He's trying to crossview.

- Honestly I find some of the most creative content on this sub

## Additional tests with alternative media
Testing other 2D images out of curiousity to see if they would generate a stereoscopic effect

### Stereo Test# 1  -   Bad Art Champion

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Bad.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Bad.png" height="280"> 
   
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Bad_CrossView.png" height="350"> 

### Stereo Test# 2  -   Simpsons

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Bart.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Bart.png" height="280"> 
    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Bart_CrossView.png" height="350"> 

### Stereo Test# 3  -   Mononoke Hime

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Hime.jpg" height="260"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Hime.png" height="260"> 
    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Hime_CrossView.png" height="350"> 

### Stereo Test# 4  -   Totoro

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Totoro.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Totoro.png" height="280"> 
   
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Totoro_CrossView.png" height="350"> 

### Stereo Test# 5  -   Totoro2

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Totoro2.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Totoro2.png" height="280"> 
    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Totoro2_CrossView.png" height="350"> 

### Stereo Test# 6  -   Wolfwalkers

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Wolfwalkers.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Wolfwalkers.png" height="280"> 
    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Wolfwalkers_CrossView.png" height="350"> 

### Stereo Test# 7  -   Alice in Wonderland

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Art_Alice.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Alice.png" height="280"> 
    
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Alice_CrossView.png" height="350"> 

### Stereo Test# 8  -   AI Painted Tree

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Paint_Tree.jpg" height="280"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Paint_Tree.png" height="280"> 

<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Paint_Tree_CrossView.png" height="350"> 


## Conclusions and Observations
The main issues seem to be the following:

- Reflective & water surfaces do not generate depth

- For best effect, use images that are wider then the object of focus to prevent cropping from the disparity

- Cartoons need good shading to generate into 3D. Strong constast appears to cause issues, however the AI Painted Tree tells us it's not so simple

