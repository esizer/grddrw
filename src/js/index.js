export class Toast {
  #duration = 300;
  #transitionClasses = ["transition", "duration-300"];

  constructor(level, message) {
    this.level = level;
    this.message = message;
  }

  #getLevelClass() {
    switch (this.level) {
      case "info":
        return ["bg-green-50", "border-green-900", "text-green-900"];
      case "warning":
        return ["bg-yellow-50", "border-yellow-900", "text-yellow-900"];
      case "danger":
        return ["bg-red-50", "border-red-900", "text-red-900"];
      default:
        return ["bg-green-50", "border-green-950", "text-green-950"];
    }
  }

  #makeWrapper() {
    const wrapper = document.createElement("div");
    wrapper.classList.add("shadow", "p-2", "border-2", "hidden");
    wrapper.classList.add(...this.#getLevelClass());
    wrapper.setAttribute("role", "alert");
    return wrapper;
  }

  #makeContent() {
    const content = document.createElement("p");
    content.textContent = this.message;
    return content;
  }

  async #enter(el) {
    el.classList.remove("hidden");

    el.classList.add(...this.#transitionClasses, "ease-out");
    el.classList.add("translate-y-1/3", "opacity-0");

    await this.#nextFrame();
    await this.#wait(this.#duration);

    el.classList.remove("translate-y-1/3", "opacity-0");
    el.classList.add("translate-y-0", "opacity-100");

    await this.#nextFrame();

    el.classList.remove(...this.#transitionClasses);
  }

  async #exit(el) {
    el.classList.add(...this.#transitionClasses, "ease-in");
    el.classList.add("translate-y-0", "opacity-100");

    await this.#nextFrame();

    el.classList.remove("translate-y-0", "opacity-100");
    el.classList.add("translate-y-1/3", "opacity-0");

    await this.#nextFrame();
    await this.#wait(this.#duration);

    el.classList.remove(...this.#transitionClasses);
    el.classList.add("hidden");

    await this.#nextFrame();

    el.remove();
  }

  #nextFrame() {
    return new Promise((resolve) => {
      requestAnimationFrame(() => {
        requestAnimationFrame(resolve);
      });
    });
  }

  #wait(dur) {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve();
      }, dur);
    });
  }

  async pop() {
    const wrapper = this.#makeWrapper();
    const content = this.#makeContent();
    wrapper.appendChild(content);

    const container = document.getElementById("toast-container");
    if (container) {
      container.appendChild(wrapper);
      await this.#enter(wrapper);
      await this.#wait(3000);
      await this.#exit(wrapper);
    }
  }
}

window.Toast = Toast;
