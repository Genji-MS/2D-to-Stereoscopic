## Stereoscopic images vs single image to stereo generation
Test results are difficult to quantify numerically. Natural stereoscopic images don't have a specific disparity between images. Actually, hyperstereo images spanning thousands of miles apart exist on the internet. Recreating an identical disparity between images is not the goal of this model. So to compare our generated stereoimages with the originals by a numerical pixel perfect accuraccy would be inaccurate.

We compare the naturally created stereoscopic images against generated images by their visual inacuracies, noting any missing or innacurate details in the generated 3D scenes. 

The generated images use only the right side of the original set and use a `x_shift_range: 0.0075` for disparity. This value is global to all of our generated images. Images were cropped using the `ImageCrop` script included in this repo. Stereoscopic images were procurred from `reddit.com/r/CrossView/`

- Vs Test# 01  -   90596

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_90596.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_90596.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_90596_CrossView.png" height="350">

    Notes:
        - At first glance it appears her shirt/abdomen is too far forward. But comparing it with the original it is absolutely correct!
        - The original image had too much disparity. 

- Vs Test# 02  -   Mirror

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mirror.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Mirror.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mirror_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 03  -   Mouth

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Mouth.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Mouth.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Mouth_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 04  -   Plant

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Plant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Plant.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Plant_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 05  -   Puddle

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Puddle.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Puddle.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Puddle_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 06  -   Xray

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Xray.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Xray.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Xray_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 07  -   Hydrant

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Hydrant.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Hydrant.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Hydrant_CrossView.png" height="350"> 

    Notes:
        - A
        - B


- Vs Test# 08  -   Dog     -   with cutouts

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Dog.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Dog.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Dog_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 09  -   Overview     -   with cutouts

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Overview.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Overview.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Overview_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 10  -   Swim     -   with cutouts

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Swim.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Swim.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Swim_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 11  -   Text

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Text.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Text.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Text_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 12  -   Art

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Test_Art.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Test_Art.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_CrossView.png" height="350"> 

    Notes:
        - A
        - B

- Vs Test# 13  -   Meme

    Original
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/3DInput/Meme.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/2D_Meme.png" height="350"> 

    Generated
    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Meme_CrossView.png" height="350"> 

    Notes:
        - A
        - B


At this point, I began to test 2D images out of curiousity to see if they would generate a stereoscopic effect

- Stereo Test# 1  -   Bad Art Champion

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Bad.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Art_Bad_CrossView.png" <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Bad_CrossView.png" height="350"> 

- Stereo Test# 2  -   Simpsons

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Bart.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Bart.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Bart_CrossView.png" height="350"> 

- Stereo Test# 3  -   Mononoke Hime

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Hime.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Hime.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Hime_CrossView.png" height="350"> 

- Stereo Test# 4  -   Totoro

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Totoro.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Totoro.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Totoro_CrossView.png" height="350"> 

- Stereo Test# 5  -   Totoro2

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Totoro2.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Totoro2.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Totoro2_CrossView.png" height="350"> 

- Stereo Test# 6  -   Wolfwalkers

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Wolfwalkers.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Wolfwalkers.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Wolfwalkers_CrossView.png" height="350"> 

- Stereo Test# 7  -   Alice in Wonderland

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Art_Alice.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Art_Alice.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Art_Alice_CrossView.png" height="350"> 

- Stereo Test# 8  -   AI Painted Tree

    - <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/2DInput/Test_Paint_Tree.jpg" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/Depth/Test_Paint_Tree.png" height="350"> <img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/2D_Test_Paint_Tree_CrossView.png" height="350"> 