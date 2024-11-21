import tkinter as tk
from tkinter import filedialog, Canvas
from PIL import Image, ImageTk

# Modes de mise à l'échelle en tant que constantes
CONTAIN = 1
COVER = 2
SCALE_DOWN = 3
FILL = 4

class ImageWidget:
    def __init__(self, root):
        self.root = root
        self.root.title("Image Display with Scaling")
        self.canvas = Canvas(root, bg="white")
        self.canvas.pack(fill=tk.BOTH, expand=True)

        self.image = None
        self.tk_image = None
        self.fit = CONTAIN  # Mode de mise à l'échelle par défaut
        self.scale = 1.0

        # Menu
        menubar = tk.Menu(root)
        file_menu = tk.Menu(menubar, tearoff=0)
        file_menu.add_command(label="Open Image", command=self.open_image)
        file_menu.add_command(label="Exit", command=root.quit)
        menubar.add_cascade(label="File", menu=file_menu)

        fit_menu = tk.Menu(menubar, tearoff=0)
        fit_menu.add_command(label="Contain", command=lambda: self.set_fit(CONTAIN))
        fit_menu.add_command(label="Cover", command=lambda: self.set_fit(COVER))
        fit_menu.add_command(label="Scale Down", command=lambda: self.set_fit(SCALE_DOWN))
        fit_menu.add_command(label="Fill", command=lambda: self.set_fit(FILL))
        menubar.add_cascade(label="Fit Mode", menu=fit_menu)

        root.config(menu=menubar)
        root.bind("<Configure>", self.resize_image)

    def open_image(self):
        file_path = filedialog.askopenfilename(
            title="Open Image",
            filetypes=[("Image Files", "*.jpg;*.jpeg;*.png;*.bmp;*.gif")]
        )
        if file_path:
            self.image = Image.open(file_path)
            self.resize_image()

    def set_fit(self, fit_mode):
        self.fit = fit_mode
        self.resize_image()

    def resize_image(self, event=None):
        if not self.image:
            return

        canvas_width = self.canvas.winfo_width()
        canvas_height = self.canvas.winfo_height()
        if canvas_width <= 1 or canvas_height <= 1:
            return

        img_width, img_height = self.image.size
        scale_x = canvas_width / img_width
        scale_y = canvas_height / img_height

        if self.fit == CONTAIN:
            scale = min(scale_x, scale_y)
        elif self.fit == COVER:
            scale = max(scale_x, scale_y)
        elif self.fit == SCALE_DOWN:
            scale = min(1.0, min(scale_x, scale_y))
        elif self.fit == FILL:
            scale_x, scale_y = scale_x, scale_y
        else:
            scale = 1.0

        if self.fit == FILL:
            new_width, new_height = int(img_width * scale_x), int(img_height * scale_y)
        else:
            new_width, new_height = int(img_width * scale), int(img_height * scale)

        resized_image = self.image.resize((new_width, new_height), Image.ANTIALIAS)
        self.tk_image = ImageTk.PhotoImage(resized_image)
        self.canvas.create_image(canvas_width // 2, canvas_height // 2, image=self.tk_image, anchor=tk.CENTER)

def main():
    root = tk.Tk()
    app = ImageWidget(root)
    root.geometry("800x600")
    root.mainloop()

if __name__ == "__main__":
    main()
