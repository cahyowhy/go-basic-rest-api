<!--suppress ALL -->
<template>
  <div class="common-editor content">
    <slot name="slot-before"/>
    <editor class="editor" :extensions="extensions" @update="onUpdate">
  
      <div class="menubar" slot="menubar" slot-scope="{ nodes, marks }">
        <div v-if="nodes && marks">
  
          <button class="menubar__button" :class="{ 'is-active': marks.bold.active() }" @click="marks.bold.command">
            <i class="mdi-format-bold mdi"></i>              
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': marks.italic.active() }" @click="marks.italic.command">
      				<i class="mdi-format-italic mdi"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': marks.strike.active() }" @click="marks.strike.command">
      			<i class="mdi mdi-format-strikethrough"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': marks.underline.active() }" @click="marks.underline.command">
      			<i class="mdi mdi-format-underline"></i>
      		</button>
  
          <button class="menubar__button" @click="marks.code.command" :class="{ 'is-active': marks.code.active() }">
      			<i class="mdi-code-tags mdi"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.paragraph.active() }" @click="nodes.paragraph.command">
      			<i class="mdi mdi-format-paragraph"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.heading.active({ level: 1 }) }" @click="nodes.heading.command({ level: 1 })">
      			H1
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.heading.active({ level: 2 }) }" @click="nodes.heading.command({ level: 2 })">
      			H2
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.heading.active({ level: 3 }) }" @click="nodes.heading.command({ level: 3 })">
      			H3
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.bullet_list.active() }" @click="nodes.bullet_list.command">
      			<i class="mdi mdi-format-list-bulleted"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.ordered_list.active() }" @click="nodes.ordered_list.command">
      			<i class="mdi mdi-format-list-numbers"></i>
      	  </button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.blockquote.active() }" @click="nodes.blockquote.command">
      			<i class="mdi-format-quote-close mdi"></i>
      		</button>
  
          <button class="menubar__button" :class="{ 'is-active': nodes.code_block.active() }" @click="nodes.code_block.command">
      			<i class="mdi-code-braces mdi"></i>
      		</button>
  
        </div>
      </div>
  
      <div class="editor__content" slot="content" slot-scope="props">
        <div class="content" v-html="newValue">
        </div>
      </div>
    </editor>
    <slot name="slot-after"/>
  </div>
</template>

<script lang="ts">
  import {
    Vue,
    Component,
    Components,
    Prop
  } from "annotation";
  import {
    Editor
  } from "tiptap";
  import {
    // Nodes
    BlockquoteNode,
    BulletListNode,
    CodeBlockNode,
    CodeBlockHighlightNode,
    HardBreakNode,
    HeadingNode,
    ImageNode,
    ListItemNode,
    OrderedListNode,
    TodoItemNode,
    TodoListNode,
  
    // Marks
    BoldMark,
    CodeMark,
    ItalicMark,
    LinkMark,
    StrikeMark,
    UnderlineMark,
  
    // General Extensions
    HistoryExtension,
    PlaceholderExtension
  } from "tiptap-extensions";
  
  @Component
  export default class CommonEditor extends Vue {
    @Prop({
      default: ""
    })
    private value: string;
  
    private extensions: Array < any > = [
      new BlockquoteNode(),
      new BulletListNode(),
      new CodeBlockNode(),
      new HardBreakNode(),
      new HeadingNode({
        maxLevel: 3
      }),
      new ImageNode(),
      new ListItemNode(),
      new OrderedListNode(),
      new TodoItemNode(),
      new TodoListNode(),
      new BoldMark(),
      new CodeMark(),
      new ItalicMark(),
      new LinkMark(),
      new StrikeMark(),
      new UnderlineMark(),
      new HistoryExtension(),
      new PlaceholderExtension()
    ];
  
    @Components
    private components() {
      return {
        Editor
      };
    }

    private get newValue() {
      return this.value;
    }

    private set newValue(value) {
      this.$emit('input', value);
    }

    private onUpdate({ getJSON, getHTML }) {
      this.$emit('update', {
        json: getJSON(),
        html: getHTML()
      });
		}
  }
</script>