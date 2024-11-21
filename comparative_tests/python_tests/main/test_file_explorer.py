import tkinter as tk
from tkinter import filedialog, messagebox
from PIL import Image, ImageTk, ImageOps, ImageDraw

class FileSelectorApp:
    def __init__(self, root):
        self.root = root
        self.root.title("Test buttons")

        self.open_btn = tk.Button(root, text="Open Image", command=self.open_image)
        self.open_btn.pack(pady=20)

        self.image_label = tk.Label(root)
        self.image_label.pack()

        self.save_btn = tk.Button(root, text="Save Image", command=self.save_image, state=tk.DISABLED)
        self.save_btn.pack(pady=20)

        self.image = None
        self.filepath = None

    def open_image(self):
        filetypes = [("JPEG files", "*.jpg"), ("All files", "*.*")]
        filepath = filedialog.askopenfilename(title="Open Image", filetypes=filetypes)
        if not filepath:
            return

        try:
            self.image = Image.open(filepath)
            self.filepath = filepath
            tk_image = ImageTk.PhotoImage(self.image)
            self.image_label.config(image=tk_image)
            self.image_label.image = tk_image
            self.save_btn.config(state=tk.NORMAL)
        except Exception as e:
            messagebox.showerror("Error", f"Failed opening image file: {e}")

    def save_image(self):
        if not self.image:
            messagebox.showerror("Error", "No file loaded, cannot save")
            return

        filepath = filedialog.asksaveasfilename(defaultextension=".jpg", filetypes=[("JPEG files", "*.jpg"), ("All files", "*.*")])
        if not filepath:
            return

        try:
            self.image.save(filepath, format="JPEG")
            messagebox.showinfo("Success", "Image saved successfully")
        except Exception as e:
            messagebox.showerror("Error", f"Failed saving image file: {e}")

def main():
    root = tk.Tk()
    app = FileSelectorApp(root)
    root.geometry("800x600")
    root.mainloop()

if __name__ == "__main__":
    main()
