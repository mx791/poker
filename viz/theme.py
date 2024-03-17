import matplotlib.pyplot as plt
import matplotlib.font_manager as font_manager

def theme(theme_color='#1F618D', second_color='#C0392B', horizontal_grid=True, back_color="#fff"):

  plt.rcParams["axes.linewidth"] = 0
  #plt.rcParams["axes.titleweight"] = "bold"
  plt.rcParams["axes.titlecolor"] = theme_color

  #plt.rcParams["axes.titlesize"] = 20
  #plt.rcParams["axes.titlelocation"] = "center"
  #plt.rcParams['axes.titlepad'] = 10
  plt.rcParams["legend.labelspacing"] = 0.5
  plt.rcParams['axes.axisbelow'] = True
  plt.rcParams['figure.facecolor'] = back_color
  plt.rcParams['axes.facecolor'] = back_color


  plt.rcParams["xtick.major.size"] = 0
  plt.rcParams["ytick.major.size"] = 0
  plt.rcParams["axes.grid"] = True
  plt.rcParams["font.family"] = 'Serif'
  plt.rcParams["axes.grid.axis"] = "y"

  font_path = './Tahoma Regular font.ttf'  # Your font path goes here
  font_manager.fontManager.addfont(font_path)
  prop = font_manager.FontProperties(fname=font_path)
  
  plt.rcParams['font.family'] = 'sans-serif'
  plt.rcParams['font.sans-serif'] = prop.get_name()

def make_chart(fc, title, subtitle, h=6, w=10, title_color='#1F618D'):
  fig, ax = plt.subplots(2, 1, gridspec_kw={'height_ratios': [1, (h-1)]}, figsize=(w, h))
  fc(ax[1])
  ax[0].text(0, 0.50, title, size=17, fontweight="bold", color=title_color)
  ax[0].text(0, 0, subtitle, size=14)
  ax[0].axis('off')
  return fig, ax