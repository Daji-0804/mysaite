import glob
import os
import shutil

#ファイル名取得
def get_file_path(folder_path):
    files = glob.glob(folder_path + "/*")
    return files

#フォルダー作成
def create_folder(folder_path):
    os.mkdir(folder_path)


read_folder = r"C:\Users\kk741\Desktop\募集要項"
output_folder = read_folder + "\copy"

#create_folder(output_folder)
files = get_file_path(read_folder)

shutil.copytree(read_folder,output_folder)
output_file_path = get_file_path(output_folder)

for file_path in output_file_path:
    os.rename(file_path,file_path.strip("_ レバテックマイページ.pdf")+".pdf")
 