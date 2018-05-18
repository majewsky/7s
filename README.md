# 7s - the Small Simple Standalone Synchronized Slide Show Server

7s can be used to present a slide show to an audience viewing on multiple
devices without having to rely on screen-sharing or video-conference apps. 7s
only provides the slide show: If the audience is not in the same room, a voice
chat app such as Mumble or Riot can be used in tandem with 7s.

## Building and running

Build with `make`, and install with `make install`. Invoke as

```
$ 7s <slide-deck-file> <listen-address>
```

The first argument is the path to the slide deck file (see below). The second
argument is the listen address where the HTTP server is started, e.g.  `:8080`
to listen on port 8080.

Once started successfully, 7s will show two URLs on stdout: The **presenter
URL** leads to the UI that the presenter uses to present the slide deck. The
**public URL** shall be shared with the audience so that they can join the
presentation. It will always show the slide currently selected by the
presenter.

## The slide deck file

The slide deck is defined in a single file, which is a stream of Markdown
documents separated by lines containing exactly three dashes. The following is
an example slide deck file with three slides:

```markdown
# My first presentation
## This is the subtitle

---

## The first content slide

* Lorem ipsum
* Dolor sit amet
  * Consectetuer
  * Adipiscing elit

---

## The second content slide

This one has an image.

![Awesome image](./images/awesome.png)

Isn't this great?
```

Paths to images are interpreted relative to the directory that contains the
slide deck file.
