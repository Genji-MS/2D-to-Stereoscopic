## 2D-to-Stereoscopic
<img src="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/Tests/StereoOutput/Art_Alice_CrossView.png" height="350"> 

Based on "3d-photo-inpainting" (Reference and links at the bottom) Instead of a movie, this repo modifies a few scripts to create a side-by-side stereoscopic image. The goal is to test the ability of 2D to 3D ML techniques. 

The novelty of this idea is to bring photos to life. Stereoscopic images have been around for ages. It's only with modern technolgy can we do things with them like full color 3D movies and VR. The basis of these technologies is to allow your individual eyes to see two images of the same scene with a slight disparity between them. Your brain will interpret these two images into 3D. While the effect can be done without the aid of technology, it takes a good deal of practice to do so comfortably.

I've always have been a fan and have practiced the technique over years. Given the chance to leverage Maching Learning, I wanted to see historic imagas, cartoons, or even a film converted into 3D.

### Exploring the models and data
The technology has come a long way, the work I am basing my research on does a great job at filling in image data. To explain futher...
  - <a href="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/EXPLORINGDATA.md">EXPLORINGDATA.md</a>

### Tests: Producing generated 3D images from naturally created 3D images
To test how good the technology is, I took a series of stereoscopic images, cropped them in half, then generated a 3D image from that. Comparing the original images with the generated ones we can see what does, and doesn't work well.
  - <a href="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/TESTS.md">TESTS.md</a>

### Parameter Tuning: Trying to get higher quality 3D
We attempt to modify some settings to see if they would produce better results
  - <a href="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/DOCUMENTATION.md">DOCUMENTATION.md</a>
  - <a href="https://github.com/Genji-MS/2D-to-Stereoscopic/blob/main/HYPERPARAMTEST.md">HYPERPARAMTEST.md</a>

### Conclusions
In conclusion. There are a few cases where images do not process well into 3D:
  - Reflective surfaces - A depth map can create the shape of glass, but it cannot create the refraction of an image through the glass. This is also true of water.
  - Vertical textures - Imagine a black and white shirt. If the black side of the shirt is outside of the frame, the model will generate the missing portion as a white shirt. The effect is minimal in an image. But would become more apparent in a video where objects enter/exit the scene or pass behind foreground objects.
  - Objects of small width - The stem of an umbrella, flowers, or individual branches. Most of these objects do not generate on a depth map.
  - Objects of similar texture. A tree, in front of a tree, in front of a tree. It becomes difficult for the network to determine the appropriate level of depth of the individual branches due to the object similarity. Trees at a distance however work well because we can observe a distinct edge of the shape.
  - Cartoons with hard lines of contrast - The backgrounds of classic animated films were often painted, and look very good in 3D. The actual cel shaded character animation will often have minimal shadows and highlights, or have solid black outlines that can generate an incorrect depth map.

## Demo
Based on the original Google Colab, we can interject a couple of scripts to generate 3D Stereoscopic images.
  - https://colab.research.google.com/drive/1ByTIIQ_dMPcaNlFFDaiuSDsQ7zO03AIo?usp=sharing


## Original License
This work is licensed under MIT License. See [LICENSE](LICENSE) for details. 
```
@inproceedings{Shih3DP20,
  author = {Shih, Meng-Li and Su, Shih-Yang and Kopf, Johannes and Huang, Jia-Bin},
  title = {3D Photography using Context-aware Layered Depth Inpainting},
  booktitle = {IEEE Conference on Computer Vision and Pattern Recognition (CVPR)},
  year = {2020}
}
```
### [[Original Paper](https://arxiv.org/abs/2004.04727)] [[Project Website](https://shihmengli.github.io/3D-Photo-Inpainting/)] [[Google Colab](https://colab.research.google.com/drive/1706ToQrkIZshRSJSHvZ1RuCiM__YX3Bz)] [[Github](https://github.com/vt-vl-lab/3d-photo-inpainting)]
